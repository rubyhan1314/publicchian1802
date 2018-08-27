function Person(){
    this.name = "wangergou";
    this.age = 24;
    this.sayHi = function(){
        return this.age;
    }
    this.sayHello = function(a,b){
        return a + b;
    }
}
var object1 = new Person();
