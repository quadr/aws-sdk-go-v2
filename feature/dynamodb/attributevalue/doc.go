// Package attributevalue provides marshaling and unmarshaling utilities to
// convert between Go types and types.AttributeValues.
//
// These utilities allow you to marshal slices, maps, structs, and scalar values
// to and from types.AttributeValue. These are useful when marshaling
// Go value tyes to types.AttributeValue for DynamoDB requests, or
// unmarshaling the types.AttributeValue back into a Go value type.
//
// AttributeValue Marshaling
//
// To marshal a Go type to a AttributeValue you can use the Marshal
// functions in the attributevalue package. There are specialized versions
// of these functions for collections of Attributevalue, such as maps and lists.
//
// The following example uses MarshalMap to convert the Record Go type to a
// types.AttributeValue type and use the value to make a PutItem API request.
//
//     type Record struct {
//         ID     string
//         URLs   []string
//     }
//
//     //...
//
//     r := Record{
//         ID:   "ABC123",
//         URLs: []string{
//             "https://example.com/first/link",
//             "https://example.com/second/url",
//         },
//     }
//     av, err := attributevalue.MarshalMap(r)
//     if err != nil {
//         panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
//     }
//
//     _, err = svc.PutItem(&dynamodb.PutItemInput{
//         TableName: aws.String(myTableName),
//         Item:      av,
//     })
//     if err != nil {
//         panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
//     }
//
// AttributeValue Unmarshaling
//
// To unmarshal a types.AttributeValue to a Go type you can use the Unmarshal
// functions in the attributevalue package. There are specialized versions
// of these functions for collections of Attributevalue, such as maps and lists.
//
// The following example will unmarshal the DynamoDB's Scan API operation. The
// Items returned by the operation will be unmarshaled into the slice of Records
// Go type.
//
//     type Record struct {
//         ID     string
//         URLs   []string
//     }
//
//     //...
//
//     var records []Record
//
//     // Use the ScanPages method to perform the scan with pagination. Use
//     // just Scan method to make the API call without pagination.
//     err := svc.ScanPages(&dynamodb.ScanInput{
//         TableName: aws.String(myTableName),
//     }, func(page *dynamodb.ScanOutput, last bool) bool {
//         recs := []Record{}
//
//         err := attributevalue.UnmarshalListOfMaps(page.Items, &recs)
//         if err != nil {
//              panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
//         }
//
//         records = append(records, recs...)
//
//         return true // keep paging
//     })
//
// The ConvertTo, ConvertToList, ConvertToMap, ConvertFrom, ConvertFromMap
// and ConvertFromList methods have been deprecated. The Marshal and Unmarshal
// functions should be used instead. The ConvertTo|From marshallers do not
// support BinarySet, NumberSet, nor StringSets, and will incorrect marshal
// binary data fields in structs as base64 strings.
//
// The Marshal and Unmarshal functions correct this behavior, and removes
// the reliance on encoding.json. `json` struct tags are still supported. In
// addition support for a new struct tag `dynamodbav` was added. Support for
// the json.Marshaler and json.Unmarshaler interfaces have been removed and
// replaced with have been replaced with attributevalue.Marshaler and
// attributevalue.Unmarshaler interfaces.
//
// `time.Time` is marshaled as `time.RFC3339Nano` format.
package attributevalue
