package models

type CreateClusterRequest struct {
	Name         string   `json:"name"`
	RoleArn      string   `json:"role_arn"`
	SubnetIds    []string `json:"subnet_ids"`
	SecurityGs   []string `json:"security_groups"`
	Version      string   `json:"version"`
}
