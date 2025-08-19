package generate

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/meopedevts/typedvo-cli/internal/codegen"
)

func generateVO(entityName string, entityFields []entityField) error {
	dir := dirName(entityName)
	className := entityName + "VO"
	file, err := os.Create(path.Join(dir, className+".kt"))
	if err != nil {
		return err
	}
	defer file.Close()

	vo := codegen.NewVOClass(className)

	imports := []string{
		"org.meopedevts.typedvo.delegates.Delegate",
		"org.meopedevts.typedvo.delegates.DelegateNotNull",
		"org.meopedevts.typedvo.delegates.DelegateBoolean",
		"br.com.sankhya.jape.vo.DynamicVO",
		"java.math.BigDecimal",
		"java.sql.Timestamp",
	}

	err = vo.AddImports(imports...)
	if err != nil {
		return err
	}

	for _, entityField := range entityFields {
		name := strings.ToLower(entityField.nomeCampo)

		genFieldDoc(entityField, vo)
		switch entityField.tipCampo {

		case "I", "F":
			vo.AddContent(fmt.Sprintf("%s var %s: BigDecimal? by Delegate()\n", codegen.Tab, name))

		case "S", "T":
			vo.AddContent(fmt.Sprintf("%s var %s: String? by Delegate()\n", codegen.Tab, name))

		case "H", "D":
			vo.AddContent(fmt.Sprintf("%s var %s: Timestamp? by Delegate()\n", codegen.Tab, name))
		}
	}

	vo.WriteToFile(file)

	return nil
}

func genFieldDoc(entityField entityField, vo *codegen.Class) {
	vo.AddContent(fmt.Sprintf("%s /*", codegen.Tab))
	vo.AddContent(fmt.Sprintf("%s * ", codegen.Tab))
	vo.AddContent(fmt.Sprintf("%s * %s - %s", codegen.Tab, entityField.nomeCampo, entityField.descrCampo))
	vo.AddContent(fmt.Sprintf("%s *", codegen.Tab))
	vo.AddContent(fmt.Sprintf("%s */", codegen.Tab))
}

func createDir(entityName string) error {
	dir := dirName(entityName)
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if err != os.ErrExist {
			return err
		}
	}
	return nil
}

func dirName(entityName string) string {
	return entityName + "Generated"
}
