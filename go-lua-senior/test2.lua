compute = function(a,b)
   print("test2.lua 参数1:",a,"参数2:",b,", 结果:",a+b)
   return a+b
end

print("======== test2.lua ========")
cb(compute)