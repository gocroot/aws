# Deplong Golang CI/CD to Amazon Web Service

This is a simple Golang Model-Controller template using [AWS Lambda Go](https://github.com/aws/aws-lambda-go/blob/main/events/README_ApiGatewayEvent.md) and mongodb.com as the database host. It is compatible with AWS Lambda CI/CD deployment.

Start here: Just [Fork this repo](https://github.com/gocroot/aws/)

## MongoDB Preparation

The first thing to do is prepare a Mongo database using this template:
1. Sign up for mongodb.com and create one instance of Data Services of mongodb.
2. Download [MongoDB Compass](https://www.mongodb.com/try/download/compass), connect with your mongo string URI from mongodb.com
3. Create database name iteung and collection reply  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/23ccddb7-bf42-42e2-baac-3d69f3a919f8)  
4. Import [this json](https://whatsauth.my.id/webhook/iteung.reply.json) into reply collection.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/7a807d96-430f-4421-95fe-1c6a528ba428)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/fd785700-7347-4f4b-b3b9-34816fc7bc53)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/ef236b4d-f8f9-42c6-91ff-f6a7d83be4fc)  

## Folder Structure

This boilerplate has several folders with different functions, such as:
* .github: GitHub Action yml configuration.
* config: all apps configuration like database, API, token.
* controller: all of the endpoint functions
* model: all of the type structs used in this app
* helper: helper folder with list of function only called by others file

## AWS Lambda CI/CD setup

Sign Up and login into aws console and go to AWS Lambda menu and follow this instruction:
1. Klik Create Function and input Function name, select Amazon Linux 2023 Runtime, select x86_64 Architecture  
   ![image](https://github.com/gocroot/aws/assets/11188109/d1728555-88ff-41e5-8b05-766e004c0c43)  
2. In Advanced settings select Enable function URL, None Aut type.
   ![image](https://github.com/gocroot/aws/assets/11188109/c600eaee-a60c-4166-b99e-da6a5b8e2fc4)  
3. Please set the environment variable in Configuration tab:  
   ![image](https://github.com/gocroot/aws/assets/11188109/f9a1e747-ab19-4498-9fe7-b7b043473a65)  
   ```sh
   MONGOSTRING=YOURMONGOSTRINGACCESS
   WAQRKEYWORD=yourkeyword
   WEBHOOKURL=https://yourappname.alwaysdata.net/whatsauth/webhook
   WEBHOOKSECRET=yoursecret
   WAPHONENUMBER=62811111
   ```
4. Go to the menu Identity and Access Management (IAM), set lambda:UpdateFunctionCode Policies, and add new users.  
   ![image](https://github.com/gocroot/aws/assets/11188109/2d489702-2aec-460b-9fe4-c319d8a6e018)  
5. Create an access key from the Security credentials tab.  
   ![image](https://github.com/gocroot/aws/assets/11188109/e24f5de5-d46d-435d-b9a6-4c2e452cc914)  
6. Go to settings>Secrets and variables>Actions in the GitHub repo. Add several Repository secrets.  
   ![image](https://github.com/gocroot/aws/assets/11188109/8e4e9c68-2beb-403f-a669-ff83b1ac04c3)  

## WhatsAuth Signup

1. Go to the [WhatsAuth signup page](https://wa.my.id/) and scan with your WhatsApp camera menu for login. 
2. Input the webhook URL(https://yourappname.alwaysdata.net/whatsauth/webhook) and your secret from the WEBHOOKSECRET setting environment on Always Data.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/e0b5cb9d-e9b3-4d04-bbd5-b03bd12293da)  
3. Follow [this instruction](https://whatsauth.my.id/docs/), in the end of the instruction you will get 30 days token using [this request](https://wa.my.id/apidocs/#/signup/signUpNewUser)
4. Save the token into MongoDB, open iteung db, create a profile collection and insert this JSON document with your 30-day token and your WhatsApp number.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/5b7144c3-3cdb-472b-8ab3-41fe86dad9cb)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/829ae88a-be59-46f2-bddc-93482d0a4999)  

   ```json
   {
     "token":"v4.public.asoiduas",
     "phonenumber":"6281111"
   }
   ```
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/06330754-9167-4bf4-a214-5d75dab7c60a)  

## Refresh Whatsapp API Token

To continue using the WhatsAuth service, we must get a new token every 3 weeks before the token expires in 30 days.
1. Open Menu Amazon EventBridge> Buses > Rules > Create Rule. Choose like screenshot.  
   ![image](https://github.com/gocroot/aws/assets/11188109/31e170af-c489-493b-bbe4-fd021157f4c8)  
2. Input for every 20 days; next, choose lambda function then set ENable state.
   ![image](https://github.com/gocroot/aws/assets/11188109/80c0869a-ae55-418c-ab7a-8f0d048bab47)  
   ![image](https://github.com/gocroot/aws/assets/11188109/11a20d9e-3bfa-436f-9549-0caf3e82f9c8)
   ![image](https://github.com/gocroot/aws/assets/11188109/828337fc-45cc-42ab-abae-ba5827b99a1d)  
4. Completing create schedule
   ![image](https://github.com/gocroot/aws/assets/11188109/94d47bb5-ad5f-46f4-a9a0-d5713ca0b06e)

## Upgrade Apps

If you want to upgrade apps, please delete (go.mod) and (go.sum) files first, then type the command in your terminal or cmd :

```sh
go mod init gocroot
go mod tidy
```

