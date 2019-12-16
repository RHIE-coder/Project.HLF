2// ExpressJS Setup
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
    if(!userExists){
        console.log('"user1" does not exist in the wallet')
        console.log('run registerUser.js')
        return;
    }
    const gateway = new Gateway()
    await gateway.connect(ccp,{wallet,identity:'user1',discovery:{enabled:false}})
    const network = await gateway.getNetwork('pretzelchannel')
    const contract = network.getContract('pretzel2')
    console.log("into the chaincode methods")
    let result;
    //inputExampleData() inputExamplePD() readExampleData() readExamplePD()
    if (fn_name === 'inputWS'){
        const u = args[0]
        const a = args[1]
        result = await contract.submitTransaction('inputWS', u,a)
    }else if(fn_name === 'inputPD'){
        const u = args[0]
        const m = args[1]
        const p = args[2]
        result = await contract.submitTransaction('inputPD', u,m,p)
    }else if(fn_name === 'readWS'){
        result = await contract.evaluateTransaction('readWS',args)
    }else if(fn_name === 'readPD'){
        const u = args[0]
        const p = args[1]
        result = await contract.evaluateTransaction('readPD', u,p)
    }else if(fn_name === 'M'){
        const a = args[0]
        const b = args[1]
        result = await contract.evaluateTransaction('M',a,b)
    }else if(fn_name === 'S'){
        result = await contract.evaluateTransaction('S',args)
    }else{
        result = "function is not found"
    }
    
    return result
}

app.post('/inputWS', async(req, res)=>{
    const username = req.body.username
    const age = req.body.age
    let args = [username, age]
    let result = await cc_call('inputWS',args)
    res.status(200).json(JSON.parse(result))
})

app.post('/inputPD', async(req, res)=>{
    const username = req.body.username
    const money = req.body.money
    const pdName = req.body.pdName
    let args = [username, money, pdName]
    let result = await cc_call('inputPD',args)
    res.status(200).json(JSON.parse(result))
})

app.post('/readWS', async (req,res)=>{
    const username = req.body.username
    let args = username
    let result = await cc_call('readWS',args)
    res.status(200).json(JSON.parse(result))
});

app.post('/readPD', async (req,res)=>{
    const username = req.body.username
    const pdName = req.body.pdName
    let args = [username,pdName]
    let result = await cc_call('readPD',args)
    res.status(200).json(JSON.parse(result))
});

app.post('/M', async (req,res)=>{
    const a = req.body.a
    const b = req.body.b
    let args = [a, b]
    let result = await cc_call('M',args)
    res.status(200).json(JSON.parse(result))
});

app.post('/S', async (req,res)=>{
    const a = req.body.a
    let args = a
    let result = await cc_call('S',args)
    res.status(200).json(JSON.parse(result))
});

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);