package cmd

import "fmt"

func Help() { 
    helpTxt :=  "1. add - create a new task \n" +
"2. done - mark a task as completed \n"+
"3. del - delete a task \n"+
"4. list - display the list of tasks "
    fmt.Println(helpTxt)
}
