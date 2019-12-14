// ExpressJS Setup
const express = require('express');
const app = express();
let bodyParser = require('body-parser');

// Hyperledger Bridge
const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', 'network' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// use static file
app.use(express.static(path.join(__dirname, 'views')));

// configure app to use body-parser
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

// main page routing
app.get('/', (req, res)=>{
    res.sendFile(__dirname + '/index.html');
})

async function cc_call(fn_name, args){
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = new FileSystemWallet(walletPath)

    const userExists = await wallet.exists('user1')
    if(!userExist){
        console.log('"user1" does not exist in the wallet')
        console.log('run registerUser.js')
        return;
    }
    const gateway = new Gateway()
    await gateway.connect(ccp,{wallet,identity:'user1',discovery:{enabled:false}})
    const network = await gateway.getNetwork('pretzelchannel')
    const contract = network.getContract('pretzel')

    let result;
    //inputExampleData() inputExamplePD() readExampleData() readExamplePD()
    if (fn_name === 'inputWS'){
        result = await contract.submitTransaction('inputExampleData', args)
    }else if(fn_name === 'inputPD'){
        result = await contract.submitTransaction(pdName, args)
    }else if(fn_name === 'readWS'){
        result = await contract.evaluateTransaction('readExampleData',args)
    }else if(fn_name === 'readPD'){
        result = await contract.evaluateTransaction(pdName, args)
    }else{
        result = "function is not found"
    }
    
}

app.post('/inputWS', async(req, res)=>{
    const username = req.body.username
    const age = req.body.age
    let args = [username, age]
    let result = cc_call('inputWS',args)
    res.status(200).json(result)
})

app.post('/inputPD/:pdName', async(req, res)=>{
    const username = req.body.username
    const money = req.body.money
    let args = [username, money,pdName]
    let result = cc_call('inputPD',args)
    res.status(200).json(result)
})

app.post('/readWS', async (req,res)=>{
    const username = req.body.username
    let args = [username]
    let result = cc_call('readWS',args)
    res.status(200).json(result)
});

app.post('/readPD/:pdName', async (req,res)=>{
    const username = req.body.username
    let args = [username,pdName]
    let result = cc_call('readPD',args)
    res.status(200).json(result)
});

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);