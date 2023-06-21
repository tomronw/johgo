package elastic

// JSON for indexes

type IndexStruct struct {
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}
type Settings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}
type Name struct {
	Type string `json:"type"`
}
type Age struct {
	Type string `json:"type"`
}
type AverageScore struct {
	Type string `json:"type"`
}
type Properties struct {
	Name         Name         `json:"name"`
	Age          Age          `json:"age"`
	AverageScore AverageScore `json:"average_score"`
}
type Mappings struct {
	Properties Properties `json:"properties"`
}
