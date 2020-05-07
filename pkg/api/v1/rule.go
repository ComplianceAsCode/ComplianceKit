package v1

type YamlRule struct {
	ID          string     `yaml:"id,omitempty"`
	Title       string     `yaml:"title,omitempty"`
	Description string     `yaml:"description,omitempty"`
	Rationale   string     `yaml:"rationale,omitempty"`
	Severity    string     `yaml:"severity,omitempty"`
	Identifiers idents     `yaml:"identifiers,omitempty"`
	References  references `yaml:"references,omitempty"`
}

type idents struct {
	CCE6 string `yaml:"cce@rhel6,omitempty"`
	CCE7 string `yaml:"cce@rhel7,omitempty"`
	CCE8 string `yaml:"cce@rhel8,omitempty"`
}

type references struct {

	// stigid@rhel6: "000145"
	Stig6 string `yaml:"stigid@rhel6,omitempty"`
	//srg@rhel6: SRG-OS-000032,SRG-OS-000037,SRG-OS-000255
	Srg6 string `yaml:"srg@rhel6,omitempty"`
	//cjis: 5.4.1.1
	Cjis string `yaml:"cjis,omitempty"`
	//cui: 3.3.1,3.3.2,3.3.6
	Cui string `yaml:"cui,omitempty"`
	//disa: 126,130,131,132,133,134
	Disa string `yaml:"disa,omitempty"`
	//hipaa: 164.308(a)(1)(ii)(D),164.308(a)(5)(ii)(C),164.310(a)(2)(iv),164.310(d)(2)(iii),164.312(b)
	Hipaa string `yaml:"hipaa,omitempty"`
	//nist: AC-2(g),AU-3,AC-17(1),AU-1(b),AU-10,AU-12(a),AU-12(c),AU-14(1),IR-5
	Nist string `yaml:"nist,omitempty"`
	//nist-csf: DE.AE-3,DE.AE-5,DE.CM-1,DE.CM-3,DE.CM-7,ID.SC-4,PR.AC-3,PR.PT-1,PR.PT-4,RS.AN-1,RS.AN-4
	Csf string `yaml:"nist-csf,omitempty"`
	//pcidss: Req-10.1
	Pci string `yaml:"pcidss,omitempty"`
	//srg: SRG-OS-000038-GPOS-00016,SRG-OS-000039-GPOS-00017,SRG-OS-000042-GPOS-00021,SRG-OS-000254-GPOS-00095,SRG-OS-000255-GPOS-00096
	Srg string `yaml:"srg,omitempty"`
	//vmmsrg: SRG-OS-000037-VMM-000150,SRG-OS-000063-VMM-000310,SRG-OS-000038-VMM-000160,SRG-OS-000039-VMM-000170,SRG-OS-000040-VMM-000180,SRG-OS-000041-VMM-000190
	VMMsrg string `yaml:"vmmsrg,omitempty"`
	//stigid@rhel7: "030000"
	Stig7 string `yaml:"stigid@rhel7,omitempty"`
	//isa-62443-2013: 'SR 1.13,SR 2.10,SR 2.11,SR 2.12,SR 2.6,SR 2.8,SR 2.9,SR 3.1,SR 3.5,SR 3.8,SR 4.1,SR 4.3,SR 5.1,SR 5.2,SR 5.3,SR 6.1,SR 6.2,SR 7.1,SR 7.6'

	//isa-62443-2009: 4.2.3.10,4.3.2.6.7,4.3.3.3.9,4.3.3.5.8,4.3.3.6.6,4.3.4.4.7,4.3.4.5.6,4.3.4.5.7,4.3.4.5.8,4.4.2.1,4.4.2.2,4.4.2.4

	//cobit5: APO10.01,APO10.03,APO10.04,APO10.05,APO11.04,APO12.06,APO13.01,BAI03.05,BAI08.02,DSS01.03,DSS01.04,DSS02.02,DSS02.04,DSS02.07,DSS03.01,DSS03.05,DSS05.02,DSS05.03,DSS05.04,DSS05.05,DSS05.07,MEA01.01,MEA01.02,MEA01.03,MEA01.04,MEA01.05,MEA02.01

	//iso27001-2013: A.11.2.6,A.12.4.1,A.12.4.2,A.12.4.3,A.12.4.4,A.12.7.1,A.13.1.1,A.13.2.1,A.14.1.3,A.14.2.7,A.15.2.1,A.15.2.2,A.16.1.4,A.16.1.5,A.16.1.7,A.6.2.1,A.6.2.2

	//cis-csc
}
