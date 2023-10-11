package handle

import (
	"fmt"
	"os"
)

func Error() {
	fmt.Println("\033[1;33m" + "Error" + "\033[0m")
	os.Exit(0)
}
