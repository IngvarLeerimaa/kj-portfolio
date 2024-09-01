const fusion = (a, b) => {
   if (dataType(a) === dataType(b)) {
     switch (dataType(a)) {
       case "array": 
       return [...a, ...b];

       case "string": 
       return a + " " + b;

       case "number": 
       return a + b;

       case "object":
         const res = { ...a };
         Object.keys(b).map((key) => { 
            res[key] = fusion(res[key], b[key]) 
         });
         return res;

       default: return b;
     }
   }
   return b;
 };
 
 const dataType = (data) => {
   switch (true) {
     case data instanceof Array: return "array";
     default: return typeof data;
   }
 };