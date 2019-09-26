local EXPRESSION = {
    __add = function (a,b)
        return Expr({add = {a,b}})
    end,
    __sub = function (a,b)
        return Expr({sub = {a,b}})
    end,
    __unm = function (a,b)
        return Expr({neg = {a,b}})
    end,
    __mul = function (a,b)
        return Expr({mul = {a,b}})
    end,
    __div = function (a,b)
        return Expr({div = {a,b}})
    end,
    __tostring = function (expr)
        if expr.name then
            return expr.name
        end
        local next = pairs(expr)
        local operator, operand = next(expr)
        print("\n\nnext", operator, operand[1], operand[2])
        return "(" .. tostring(operand[1]) .. " " .. tostring(operator) .. " " .. tostring(operand[2]) .. ")"
    end
}

function Expr(expr)
    setmetatable(expr, EXPRESSION)
    return expr
end

function Variable(name)
    local var = {}
    var.name = name
    setmetatable(var, EXPRESSION)
    print("new var", var.name)
    return var
end

setmetatable(_G, {
    __index = function (tbl, key)
        return Variable(key)
    end
})

function min (fns)
    local fn = fns[1]
    print("minimize", tostring(fn))
    printTbl(fn)
end

function printTbl(tbl, indent)
    indent = indent or 0
    for i,v in pairs(tbl) do
        if type(v) == "table" then
            print(string.rep("\t", indent), i)
            printTbl(v, indent+1)
        else
            print(string.rep("\t", indent), i,v)
        end
    end
end

min {-2 * x - 3 * y - 4 * z}

--[[ subjectTo {
    3 * x + 2 * y + z <= 10, 
    2 * x + 5 * y + 3 * z <= 15, 
    x, y, z >= 0
} ]]
