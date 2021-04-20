# medium-build-custom-plugins
 Repository containing all the custom plugins used by the Medium article on Direktiv custom plugins

## Get-Tweets

Contains the Go code and the Dockerfile ready for use in Direktiv. The code takes a simple JSON structured data as input:

Input:
```json
{
   "username" : "Twitter"
 }
```
 
Output:
```json
[
   {
      "Hashtags" : null,
      "Videos" : null,
      "Retweets" : 21394,
      "IsPin" : false,
      "Photos" : null,
      "Replies" : 2792,
      "Text" : "RT for tired, like for very tired",
      "IsReply" : false,
      "Likes" : 118543,
      "IsRetweet" : false,
      "Retweet" : {
         "UserID" : "",
         "TimeParsed" : "0001-01-01T00:00:00Z",
         "Timestamp" : 0,
         "ID" : "",
         "Username" : ""
      },
      "Error" : null,
      "Username" : "Twitter",
      "HTML" : "RT for tired, like for very tired",
      "Timestamp" : 1617925686,
      "ID" : "1380306486962782208",
      "URLs" : null,
      "TimeParsed" : "2021-04-08T23:48:06Z",
      "IsQuoted" : false,
      "PermanentURL" : "https://twitter.com/Twitter/status/1380306486962782208",
      "UserID" : "783214"
   },
   {
      "Likes" : 67902,
      "IsRetweet" : false,
      "Retweet" : {
         "UserID" : "",
         "TimeParsed" : "0001-01-01T00:00:00Z",
         "Username" : "",
         "Timestamp" : 0,
         "ID" : ""
      },
      "Error" : null,
      "Username" : "Twitter",
      "HTML" : "do they like you or do they just like your Tweet",
      "Hashtags" : null,
      "Videos" : null,
      "Retweets" : 8460,
      "IsPin" : false,
      "Photos" : null,
      "Replies" : 5696,
      "Text" : "do they like you or do they just like your Tweet",
      "IsReply" : false,
      "PermanentURL" : "https://twitter.com/Twitter/status/1379490826171072514",
      "UserID" : "783214",
      "Timestamp" : 1617731217,
      "ID" : "1379490826171072514",
      "URLs" : null,
      "TimeParsed" : "2021-04-06T17:46:57Z",
      "IsQuoted" : false
   }
]   
```
