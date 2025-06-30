package dto

// Relation DTO
type RelationDTO struct {
	Object   string `json:"o"`
	Parent   string `json:"p"`
	Metadata string `json:"m"`
}

// Add new object-object relation
type AddRelationRq RelationDTO
type AddRelationRp struct{}

// Get all relations, where reuquested object is parent
type GetChildRelationsRq struct {
	ParentObject string `json:"p"`
}
type GetChildRelationsRp struct {
	Relations []RelationDTO `json:"r"`
}

// Get relation tree including parent
type GetChildRelationsTreeRq struct {
	ParentObject string `json:"p"`
	TreeMaxDepth int    `json:"d"`
}
type TreeRelationDTO struct {
	Id       string `json:"i"`
	Metadata string `json:"m"`

	Object *TreeRelationDTO `json:"o"`
	Parent *TreeRelationDTO `json:"p"`
}

type GetChildRelationsTreeRp struct {
	Head TreeRelationDTO `json:"h"`
}

// Delete relation between parent and object
type DeleteRelationRq struct {
	Parent string `json:"p"`
	Object string `json:"o"`
}
type DeleteRelationRp struct{}

// Update relation metadata
type UpdateRelationMetadataRq RelationDTO
type UpdateRelationMetadataRp struct{}
