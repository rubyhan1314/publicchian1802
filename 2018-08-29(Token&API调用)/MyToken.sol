pragma solidity ^0.4.24;

//step1:定义接口
contract EIP20Interface{
// 	//获取token的名字，比如"KongYiXueYuan","RubyToken","BitCoin"
//     function name() view returns (string name);
//     //获取token的简称，比如"KTXYB","RTH","BTH"
//     function symbol() view returns (string symbol);
//     //获取小数位，以太坊的decimals为18
//     function decimals() view returns (uint8 decimals);
    //获取token发布的总量，比如 EOS 10亿
    
     uint256 public totalSupply;
    // function totalSupply() view returns (uint256 totalSupply);
    //获取_owner地址的余额
    function balanceOf(address _owner) view returns (uint256 balance);
    //转账，调用者向_to，转账，_value个token
    function transfer(address _to, uint256 _value) returns (bool success);
    //转账，从_from,转账到_to，_value个token
    /*
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success);

    //调用方， 允许_spender,从自己地址上转出_value，个token
    function approve(address _spender, uint256 _value) returns (bool success);
    //自己_owner查询_spender这个地址可以转走自己的多少个token
    function allowance(address _owner, address _spender) view returns (uint256 remaining);
*/

    //转账的时候必须要调用的事件
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    //成功执行approve方法后调用的事件
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}

//step2:定义实现类
contract MyToken is EIP20Interface{
    
    //step3:定义状态变量，接收创建合约时，传入的数据：代币的名称，简称，小数位，总量
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    
    //step4：提供构造函数
    function MyToken(string _name,string _symbol,uint8 _decimals,uint256 _totalSupply){
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        totalSupply = _totalSupply;
        balances[msg.sender] = _totalSupply;
    }
    
    //step5:部署合约后，一次性发布给合约部署者。
    //其他人转账后，拥有者地址，有多少个币，map[address]-->uint256
    mapping (address=>uint256) balances;
    
    
    
    //step6:获取_owner地址的余额
    function balanceOf(address _owner) view returns (uint256 balance){
        return balances[_owner];
    }
    
    //step7:转账，调用者向_to，转账，_value个token
    function transfer(address _to, uint256 _value) returns (bool success){
        
        require(_value > 0 && balances[msg.sender] > _value && balances[_to] + _value > balances[_to]);//条件不满足，抛出异常，满足就继续向下执行
        
        balances[_to] += _value;
        balances[msg.sender] -= _value;
        
        //触发事件
        Transfer(msg.sender, _to, _value);
        return true;

        
    }
    
    //转账，从_from,转账到_to，_value个token
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success){
        //获取_from地址允许_to地址转走的金额
        uint256 allowan = allowances[_from][_to];
        //_from账户余额大于等于要转走的金额
        //要转走的金额，小于等于允许转走金额
        //从哪转账：to是调用者
    
        require(_value > 0 && balances[_from] >= _value && allowan >= _value && _to == msg.sender && balances[_to]+_value > balances[_to]);
        
        
        allowances[_from][_to] -= _value;
        
        balances[_from] -= _value;
        balances[_to] += _value;
        
        Transfer(_from,_to,_value);
        return true;
        
        
    }
    //记录每个地址允许别人从自己账户转走金额的map
    mapping(address=>mapping(address=>uint256) )allowances;
    //step8:调用方， 允许_spender,从自己地址上转出_value，个token
    function approve(address _spender, uint256 _value) returns (bool success){
          /*
        当前调用者，允许_spender账户从地址地址上
        如果是多次调用，
        */

        allowances[msg.sender][_spender] = _value;
        //触发事件
        Approval(msg.sender,_spender,_value);

        return true;
    }
    //step9:自己_owner查询_spender这个地址可以转走自己的多少个token
    function allowance(address _owner, address _spender) view returns (uint256 remaining){
        return allowances[_owner][_spender];
    }
    
}