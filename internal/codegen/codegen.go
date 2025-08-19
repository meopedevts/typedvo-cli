package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

var (
	errNoImportsFound           = errors.New("nenhum import encontrado")
	errConstructorAlreadyExists = errors.New("constructor já existente")
)

var Tab = "    "

type Class struct {
	suffix    string
	className string
	imports   []string

	buffer bytes.Buffer

	hasImports     bool
	hasConstructor bool
}

func newDefaultClass(className string) *Class {
	c := &Class{
		className: className,
	}
	c.addPackage()

	return c
}

func NewVOClass(className string) *Class {
	class := newDefaultClass(className)
	class.suffix = "VO"

	return class
}

func NewDAOClass(className string) *Class {
	class := newDefaultClass(className)
	class.suffix = "DAO"

	return class
}

func (c *Class) addPackage() {
	c.buffer.WriteString("package org.meopedevts.typedvo.thankyouforuseme")
	c.NewDoubleLine()
}

func (c *Class) addConstructor() {
	c.buffer.WriteString(fmt.Sprintf("class %s(vo: DynamicVO): DynamicVO by vo {\n", c.className))
	c.hasConstructor = true
}

func (c *Class) AddImports(imports ...string) error {
	if c.hasConstructor {
		return fmt.Errorf("%w: não é possível inserir mais imports", errConstructorAlreadyExists)
	}

	for _, imp := range imports {
		c.buffer.WriteString(fmt.Sprintf("import %s\n", imp))
	}
	c.NewLine()
	c.hasImports = true

	return nil
}

func (c *Class) AddContent(content string) error {
	if !c.hasImports {
		return fmt.Errorf("%w: %s", errNoImportsFound, c.className)
	}

	if !c.hasConstructor {
		c.addConstructor()
	}

	c.buffer.WriteString(content)
	c.NewLine()

	return nil
}

func (c *Class) WriteToFile(file *os.File) {
	if !c.hasConstructor {
		c.addConstructor()
	}
	c.buffer.WriteString("}")
	c.buffer.WriteTo(file)
}

// Utils

func (c *Class) NewLine() {
	c.buffer.WriteString("\n")
}

func (c *Class) NewDoubleLine() {
	c.buffer.WriteString("\n\n")
}
