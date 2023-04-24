function Person(name, age, sex) {
    this.name = name;
    this.age = age;
    this.sex = sex;
    this.getName = function() {
        return this.name;
    };
}
var rand = new Person("Rand McKinnon", 33, "M");

console.log(rand.getName())