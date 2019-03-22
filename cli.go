package main 

import (
	"log"
	"os"

    "github.com/urfave/cli"
)
func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Lets you query IP's, CNAMES, MXRecords and Name Services "

	myFlags := []cli.NewApp()
	cli.StringFlag{ 
			Name: "host"
			Value: "tutorialedge.net"
	},

	}
	app.Commands = []cli.Command{
			Name: "ns",
			Usage: "Looks up the NameServers for the particular host",
			Flags:myFlags,
			Action: func( c *cli.Context)error{
				ns,err := net.LookupNS(c.String("url"))
				if err != nil{
					return err
				}
				for i =0 i< len(ns); i++{
					fmt.Println(ns[i].Host)
				}
				return nil
			},

	},
	{
		Name: "cname",
		Usage: "Lookup the CNAME for the particular host"
		Flags: myFlags,
		Action : func( c  *cli.Context) error{
			cname, err := net.LookupCNAME(c.String("host"))
			if err != nil{
					fmt.Println(err)
			}
			fmt.Println(cname)
			return nil
		},
	},
{
	Name: "ip",
	Usage: "Looksup the IP Addresses for the particular host",
	Flags: myFlags,
	Action: func( c *cli.Context) error {
		ip,err := netLookupIP(c.String("host"))
		if err != nil{
			fmt.println(err)
		}
		for i:= 0; i<len(p);i++{
			fmt.Println(ns[i].Host)
		}
		return nil
	},
},
{
	Name: "cname",
	Usage: "Looks up the CNAME fortjhe particular host "
	Flags: myFlags,
	Action: func( c *cli.Context) error {
		cname, err := net.LookupCNAME(c.String("host"))
		if err != nil{
			fmt.println(err)
		}
		fmt.Println(cname)
		return nil
	},
},
{ 
		Name: "ip",
		Usage "Lookus up the IP addresses for the particular host ",
		Flags: myFlags,
		Action: func(c *cli.Context) error{
			ip, err := net.LookupIP(c.String("host"))
			if err != nil{
				fmt.Println(err)
			}
			for i:= 0; i< len(ip); i++ {
				fmt.Println(ip[i])
			}
			return nil
		},
},
{ 
	Name:"mx",
	Usage: "Looks up the MX Records for the particular host "
	Flags: myFlags,
	Action: func(c *cli.Context) error {
		mx, err := net.LookupMX(c.String("host"))
		if err := nil{
				fmt.Println(err)
		}
		for i  := 0;i<len(mx); i++{
				fmt.Println(mx[i], Host, mx[i].Pref)
		}
		return nil
	},
},
}

err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}