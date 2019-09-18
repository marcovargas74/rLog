/********************************************************************
 *    Descrição: Funções de uso Geral
 *
 *    Autor: Marco Antonio Vargas
 *
 *    Data:02/07/2019
 *******************************************************************/

package rlog

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//CertificatePathDefault é o caminho default do certificado
const (
	versionPackage = "2019-09-18"
	//Level of Logs
	Emerg   = syslog.LOG_EMERG
	Alert   = syslog.LOG_ALERT
	Crit    = syslog.LOG_CRIT
	Err     = syslog.LOG_ERR
	Warning = syslog.LOG_WARNING
	Notice  = syslog.LOG_NOTICE
	Info    = syslog.LOG_INFO
	Debug   = syslog.LOG_DEBUG
	//Facility
	Local0 = syslog.LOG_LOCAL0
	Local1 = syslog.LOG_LOCAL1
	Local2 = syslog.LOG_LOCAL2
	Local3 = syslog.LOG_LOCAL3
	Local4 = syslog.LOG_LOCAL4
	Local5 = syslog.LOG_LOCAL5
	Local6 = syslog.LOG_LOCAL6
	Local7 = syslog.LOG_LOCAL7
)

//GLOBAL VAR

//AppLog Variavel usado no syslog
var AppLog *syslog.Writer

//var AppLog io.Writer

//AppLevel onde armazenado o nivel do syslog
var AppLevel syslog.Priority

//AppLogProg se Log Programado True
var AppLogProg bool

//AppLogFprintSyslog imprime print F
var AppLogFprintSyslog = false

//Clear Limpa a Tela
func Clear() {
	//fmt.Println("\033[2J")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//GetVersion Get the number version of packet
func GetVersion() string {
	return versionPackage
}

//SetPrintLocal Set to true is want to print local message like fmt.printf(..)
//default is false
func SetPrintLocal(isPrint bool) {
	AppLogFprintSyslog = isPrint
}

// ThisFunction return a string containing the file name, function name
// and the line number of a specified entry on the call stack
//func ThisFunction(depthList ...int) string {
func ThisFunction() string {
	/*var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)*/
	function, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s->%s(%d)", chopPath(file), chopPath(runtime.FuncForPC(function).Name()), line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	//if i == -1 {
	//return original
	//}
	return original[i+1:]
}

//============================= LOGGIN  ==================================

/*Formas de Imprimir Logs
 *fmt.Fprintf(AppLog, "%s sys/Log  Iniciado\n", ThisFunction())
 *sysLog.Debug("And this is a daemon emergency with demotag.")
 *log.Printf("%s sys/Log  Iniciado\n", ThisFunction())
 */

//StartLogger Inicia Login da aplicação
//isProg true to show message
//level priority of message
//serverIP where the message will be sent
func StartLogger(isProg bool, level syslog.Priority, serverIP string) {
	var err error

	AppLogProg = isProg
	AppLevel = level
	//AppLevel = syslog.LOG_DEBUG | syslog.LOG_LOCAL7

	if isProg == false {
		return
	}

	//AppLog, err = syslog.Dial("udp", serverIP, level, "app_simula")
	AppLog, err = syslog.Dial("udp", serverIP, level, "appLog")
	if err != nil {
		log.Fatal(err)
	}
	defer AppLog.Close()
	AppSyslog(syslog.LOG_INFO, "%s sys/Log  Iniciado\n", ThisFunction())
}

//LoggerClose Finish the Logger
func LoggerClose() {
	AppSyslog(syslog.LOG_INFO, "%s {LOG_FINISH}\n", ThisFunction())
	AppLog.Close()
}

/* App_syslog
  *Aplicar o Padrao enm todas as Mensagens de LOG
  * Sintaxe da Mensagem de LOG
  * (funcao)    = Entre parenteses: Nome Da Funcao
  * {texto}    = Entre chaves: Texto qualquer
  * <variavel> = Entre Couchetes NOme da Variavel
  * [%x x] = Entre Couchetes Valor da Variavel

	icip_syslog( LOG_INFO ,"%s:(FUNC){Este eh umLOG de teste}<argument_1>[%d]:",THIS_FILE, arg1);
  int app_syslog(int syslogpri, char *format, ...)
*/

//AppSyslog mensagem de log Padrao
func AppSyslog(syslogpri syslog.Priority, format string, a ...interface{}) {

	if syslogpri > AppLevel {
		return
	}

	//Imprime informacoes na tela como prints
	info := fmt.Sprintf(format, a...)
	if AppLogFprintSyslog {
		log.Printf("%v", info)
	}

	//SYSLOG INFO
	if syslogpri == syslog.LOG_INFO {
		AppLog.Info(info)
		return
	}

	//SYSLOG ERRO
	if syslogpri == syslog.LOG_ERR {
		AppLog.Err(info)
		return
	}

	//Nivel Debug
	fmt.Fprintf(AppLog, format, a...)
}

//------------------ FIM ARQUIVO GO --------------------------------------
