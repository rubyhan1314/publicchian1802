pragma solidity ^0.4.24;

contract Voting{
    
    //step1:创建候选人
    struct Candidate{
        uint voteCount;
        string name;
    }
    
    //step2:voter
    struct Voter{
        bool voted;
    }
    
    //step3:voters
    mapping(address=>Voter) public voters;
    
    //step4:candidates
    Candidate[] public candidates;
    
    //step5:构造函数中，初始化候选人
    function Voting()public{
        candidates.push(Candidate(0,"wangergou")); //index:0
        candidates.push(Candidate(0,"lixiaohua")); //index:1
    }
    
    //step6:投票方法
    function vote_candidate(uint index){
        if (voters[msg.sender].voted || index > candidates.length){
            return;
        } 
        candidates[index].voteCount += 1;
        voters[msg.sender].voted = true;
    }
    
    //step7:获取结果
    function getResult() constant public returns(string,uint,string,uint){
        return (candidates[0].name,candidates[0].voteCount,candidates[1].name,candidates[1].voteCount);
    }
}