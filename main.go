package main

import (
        "encoding/xml"
        "fmt"
        "github.com/projectdiscovery/goflags"
        "github.com/projectdiscovery/gologger"
        "io/ioutil"
        "log"
        "os"
)

type options struct {
        silent bool
        file   string
}

func main() {
        opt := &options{}

        flagSet := goflags.NewFlagSet()
        flagSet.SetDescription(`
██╗  ██╗███╗   ███╗██╗                         
╚██╗██╔╝████╗ ████║██║                         
 ╚███╔╝ ██╔████╔██║██║                         
 ██╔██╗ ██║╚██╔╝██║██║                         
██╔╝ ██╗██║ ╚═╝ ██║███████╗                    
╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝                    
                                               
████████╗██╗  ██╗██╗███╗   ██╗ ██████╗ ███████╗
╚══██╔══╝██║  ██║██║████╗  ██║██╔════╝ ██╔════╝
   ██║   ███████║██║██╔██╗ ██║██║  ███╗███████╗
   ██║   ██╔══██║██║██║╚██╗██║██║   ██║╚════██║
   ██║   ██║  ██║██║██║ ╚████║╚██████╔╝███████║
   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝
                     Created by Zero To Hero :)`)
        flagSet.StringVar(&opt.file, "file", "", "xml file to be parsed")
        flagSet.BoolVar(&opt.silent, "silent", false, "silent mode")
        if err := flagSet.Parse(); err != nil {
                log.Fatalf("Could not parse flags: %s\n", err)
        }
        if opt.silent != true {
                banner()
        }

        xmlFile, err := os.Open(opt.file)
        if err != nil {
                log.Fatalln(err)
        }
        defer xmlFile.Close()

        val, err := ioutil.ReadAll(xmlFile)
        if err != nil {
                log.Fatalln(err)
        }
        var nmaprun NmapRun
        if err := xml.Unmarshal(val, &nmaprun); err != nil {
                log.Fatalln(err)
        }

        for _, host := range nmaprun.Hosts {
                if host.Hostnames.Name != "" {
                        fmt.Println(host.Hostnames.Name)
                }
        }
}

type NmapRun struct {
        Scanner          string    `xml:"scanner,attr"`
        Args             string    `xml:"args,attr"`
        Start            string    `xml:"start,attr"`
        StartStr         string    `xml:"startstr,attr"`
        Version          string    `xml:"version,attr"`
        XmlOutputVersion string    `xml:"xmloutputversion,attr"`
        ScanInfo         ScanInfo  `xml:"scaninfo"`
        Verbose          Verbose   `xml:"verbose"`
        Debugging        Debugging `xml:"debugging"`
        Hosts            []Host    `xml:"host"`
}

type ScanInfo struct {
        Type        string `xml:"type,attr"`
        Protocol    string `xml:"protocol,attr"`
        NumServices string `xml:"numservices,attr"`
        Services    string `xml:"services,attr"`
}

type Verbose struct {
        Level string `xml:"level,attr"`
}

type Debugging struct {
        Level string `xml:"level,attr"`
}

type Host struct {
        StartTime string    `xml:"starttime,attr"`
        EndTime   string    `xml:"endtime,attr"`
        Status    Status    `xml:"status"`
        Address   Address   `xml:"address"`
        Hostnames Hostnames `xml:"hostnames"`
        Ports     Ports     `xml:"ports"`
}

type Status struct {
        State     string `xml:"state,attr"`
        Reason    string `xml:"reason,attr"`
        ReasonTtl string `xml:"reason_ttl,attr"`
}

type Address struct {
        Addr     string `xml:"addr,attr"`
        AddrType string `xml:"addrtype,attr"`
}

type Hostnames struct {
        Name string `xml:"name,attr"`
}

type Ports struct {
        ExtraPorts ExtraPorts `xml:"extraports"`
}

type ExtraPorts struct {
        State        string       `xml:"state,attr"`
        Count        string       `xml:"count,attr"`
        ExtraReasons ExtraReasons `xml:"extrareasons"`
}

type ExtraReasons struct {
        Reason string `xml:"reason,attr"`
        Count  string `xml:"count,attr"`
        Proto  string `xml:"proto,attr"`
        Ports  string `xml:"ports,attr"`
}

func banner() {
        gologger.Print().Msgf(`
██╗  ██╗███╗   ███╗██╗                         
╚██╗██╔╝████╗ ████║██║                         
 ╚███╔╝ ██╔████╔██║██║                         
 ██╔██╗ ██║╚██╔╝██║██║                         
██╔╝ ██╗██║ ╚═╝ ██║███████╗                    
╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝                    
                                               
████████╗██╗  ██╗██╗███╗   ██╗ ██████╗ ███████╗
╚══██╔══╝██║  ██║██║████╗  ██║██╔════╝ ██╔════╝
   ██║   ███████║██║██╔██╗ ██║██║  ███╗███████╗
   ██║   ██╔══██║██║██║╚██╗██║██║   ██║╚════██║
   ██║   ██║  ██║██║██║ ╚████║╚██████╔╝███████║
   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝
`)
        gologger.Print().Msgf("   Created by Zero To Hero :)\n\n")
}
