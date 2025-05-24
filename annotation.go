package entvis

const VisibilityAnnotationName = "Visibility"
const RolesAnnotationName = "Roles"

type VisibilityAnnotation []string

func (v VisibilityAnnotation) Name() string {
	return VisibilityAnnotationName
}

func Visibility(roles ...string) VisibilityAnnotation {
	return VisibilityAnnotation(roles)
}
