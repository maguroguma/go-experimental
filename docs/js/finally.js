function deferExample() {
  console.log("Start");

  try {
    if (Math.random() > 0.5) {
      console.log("Early return 1");
      return "Result 1";
    }

    console.log("Early return 2");
    return "Result 2";
  } finally {
    console.log("Cleanup");
  }
}

function other() {
  console.log("Start");

  try {
    console.log("in try");
    if (Math.random() > 0.5) {
      console.log("Early return 1");
      return;
    }
  } finally {
    console.log("in finally");
  }

  console.log("End");
}

// fn に defer という関数を渡しつつ、代理で fn を実行し結果を得る
function withDefer(fn) {
  const deferred = []; // defer スタック
  const defer = (callback) => deferred.push(callback);

  const result = fn(defer); // 本来実行したい処理を実行する

  while (deferred.length) {
    deferred.pop()(); // LIFO で実行
  }

  return result;
}

// 例: 複数の return がある関数
function deferExample() {
  return withDefer((defer) => {
    console.log("Start");

    defer(() => console.log("Cleanup 1st"));

    if (Math.random() > 0.5) {
      console.log("Early return 1");
      return "Result 1";
    }

    defer(() => console.log("Cleanup 2nd"));

    console.log("Early return 2");
    return "Result 2";
  });
}

// console.log(example());
// other();

console.log(deferExample());
