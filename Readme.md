### Generate Documentation Swagger
swag init --generalInfo cmd/groot/main.go --output cmd/groot/docs
<!-- 
// @Param uuid path string true "The UUID of a thing"
//
// Explanation
// Name: uuid
// Where: path/body/query
// Type: string/int/{object}
// Required: true/false
// Description: "The UUID of a thing"
 -->