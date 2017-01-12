//Github response structure
package messages;

message List{

  message ghList {
    string name =1;
    string stargazers_count =2;
  }

  string token =1;
  repeated ghList =2;
}
