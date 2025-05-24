package entvis

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed views.tmpl
var viewsTemplate string

type ViewExtension struct {
	entc.DefaultExtension
}

// func (ViewExtension) Annotations() []entc.Annotation{}
// func (ViewExtension) Options() []entc.Option

func NewViewExtension() entc.Extension {
	return &ViewExtension{}
}

func (ViewExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		GenerateRolesForModels(),
	}
}

func (ViewExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("views").Parse(viewsTemplate)),
	}
}

type ModelRole struct {
	PrivateName string
	PublicName  string

	Fields []*gen.Field
}

func GenerateRolesForModels() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			allRoles := make(map[string]*ModelRole)

			for _, node := range g.Nodes {
				if node.IsView() {
					continue
				}

				var modelMappedRoles = make(map[string]*ModelRole)

				for _, field := range node.Fields {
					iVisibility, found := field.Annotations[VisibilityAnnotationName]
					if !found {
						continue
					}

					visibility, ok := iVisibility.([]any)
					if !ok {
						return fmt.Errorf("visibility annotation is not of type VisibilityAnnotation, got: %T", iVisibility)
					}

					for _, iRoleName := range visibility {
						publicRoleName, ok := iRoleName.(string)
						if !ok {
							return fmt.Errorf("visibility annotation is not of type VisibilityAnnotation, got: []%T", iRoleName)
						}

						publicRoleName = cases.Title(language.Und, cases.NoLower).String(publicRoleName)
						modelRole, found := modelMappedRoles[publicRoleName]
						if !found {
							modelRole = &ModelRole{
								PublicName:  publicRoleName,
								PrivateName: strings.ToLower(publicRoleName),
							}
							modelMappedRoles[publicRoleName] = modelRole
							allRoles[publicRoleName] = modelRole
						}

						modelRole.Fields = append(modelRole.Fields, field)
					}
				}

				var roles []*ModelRole
				for _, roleKey := range slices.Sorted(maps.Keys(modelMappedRoles)) {
					roles = append(roles, modelMappedRoles[roleKey])
				}

				node.Annotations.Set(RolesAnnotationName, roles)
			}

			var roles []*ModelRole
			for _, role := range slices.Sorted(maps.Keys(allRoles)) {
				roles = append(roles, allRoles[role])
			}

			g.Annotations.Set(RolesAnnotationName, roles)

			return next.Generate(g)
		})
	}
}
