package nuclei

import (
	"github.com/niudaii/zpscan/internal/utils"
	"github.com/projectdiscovery/nuclei/v2/pkg/templates"
	"gopkg.in/yaml.v3"
	"strings"
)

// LoadAllExp 加载全部exp
func LoadAllExp(pocDir string) (exps []*Template, err error) {
	var pocPathList []string
	pocPathList, err = utils.GetAllFile(pocDir)
	if err != nil {
		return
	}
	for _, pocPath := range pocPathList {
		if !strings.HasSuffix(pocPath, "-exp.yaml") {
			continue
		}
		//var exp *Template
		//exp, err = templates.Parse(pocPath, nil, ExecuterOptions)
		//if err != nil {
		//	gologger.Error().Msgf("ParsePocFile() %v err, %v", pocPath, err)
		//	continue
		//}
		var data []byte
		data, err = utils.ReadFile(pocPath)
		if err != nil {
			return
		}
		template := &templates.Template{}
		if err = yaml.Unmarshal(data, template); err != nil {
			return
		}
		template.Path = pocPath
		exps = append(exps, template)
	}
	return
}
