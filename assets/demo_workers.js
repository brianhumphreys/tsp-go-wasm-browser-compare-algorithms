importScripts("/assets/wasm_exec.js");

// import "/assets/wasm_exec.js";
// import { divide } from "/assets/main.go";

// const result = await divide(6, 2);
// console.log("holy shit it worked")
// console.log(result); // 3
// var i = 0;

function createWasmArray (array) {
    const arrayMap = {};
    for (let i =  0; i < array.length; i++) {
        arrayMap[`${i}`] = array[i];
    }
    arrayMap['length'] = array.length;
    console.log(arrayMap);
    return arrayMap;
}

function timedCount() {

    // console.log(self.global.formatJSON);
    // const testJson = '{"website":"golangbot.com", "tutorials": {"string":"https://golangbot.com/strings/"}}'
    // console.log(self.global.formatJSON({json: testJson}));
    // console.log(self.global.distance([10, 10], [13, 14]));

    const array = [{x: 10, y: 10}, {x: 13, y: 14}, {x: 45, y: 17}, {x: 18, y: 5}, {x: 8, y: 18}]
    // console.log(createWasmArray(array));
    // for (i, item in enumerate(array)) {
    //     mapped[i] = item
    //     console.log(item)
    // }
    // console.log(mapped)
    console.log(self.global.cost(createWasmArray(array)));
    
    // postMessage(i);
    setTimeout(() => timedCount(), 500);
}

const go = new Go();
WebAssembly.instantiateStreaming(
    fetch("../test.wasm"),
    go.importObject
  ).then((result) => {
      console.log("faclk")
      go.run(result.instance);
      timedCount();
      
  });




