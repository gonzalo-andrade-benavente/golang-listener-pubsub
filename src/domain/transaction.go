package domain

/*
2023/11/12 10:19:01 Event received: ID=1dd50998-67a2-42bd-afdd-cb053b6b0a69,
Payload={"eventId":"acf81161-6151-429e-8602-9ce9cee66b3c","eventType":"pixCreated"
,"entityId":"149176170_1699454112636","entityType":"pix","timestamp":"1699454112","datetime":"2023-11-08T14:35:12","version":"1.0.0",
"country":"CL","commerce":"Tottus","channel":"WEB","cmref":"TGFA-WMOS","domain":"SCHE","capability":"whmg","mimeType":"application/json","data":{"transactionNumber":"149176170_1699454112636"}}
*/

type Transaction struct {
	EventId    string `json:"eventId"`
	EventyType string `json:"eventType"`
	EntityId   string `json:"entityId"`
	EntityType string `json:"entityType"`
	Timestamp  string `json:"timestamp"`
	Datetime   string `json:"datetime"`
	Version    string `json:"version"`
	Country    string `json:"country"`
	Commerce   string `json:"commerce"`
	Channel    string `json:"channel"`
	Cmref      string `json:"cmref"`
	Domain     string `json:"domain"`
	Capability string `json:"capability"`
	MimeType   string `mimeType:"eventId"`
	Data       struct {
		TransactionNumber string `json:"transactionNumber"`
	} `json:"data"`
}
