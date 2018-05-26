// 修改	echo程序，使其能够打印	os.Args[0]	，即被执行命令本身的名字
package	main

import	(	
	"fmt"	
	"os" 
)
func main()	{	
	s,	sep	:=	"",	""				
	for	_,	arg	:=	range os.Args[1:] {
			s	+=	sep	+	arg								
			sep	=	"	"				
	}
	fmt.Println(os.Args[0])		
	fmt.Println(s)
}
