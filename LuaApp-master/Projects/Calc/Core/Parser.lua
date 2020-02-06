--
-- Created by IntelliJ IDEA.
-- User: Камиль
-- Date: 30.01.2020
-- Time: 15:03
-- To change this template use File | Settings | File Templates.
--
Parser = {}

Token = {}
Token.NUMBER = "NUMBER"
Token.PLUS = "PLUS"
Token.MINUS = "MINUS"
Token.SUB = "SUB"
Token.DIV = "DIV"

function Token:new(type, value)
    local obj = {}
    obj.Type = type
    obj.Value = value

    function obj:getType()
        return obj.Type
    end

    function obj:getValue()
        return obj.Value
    end

    setmetatable(obj, self)
    self.__index = self; return obj
end

local function getNextChar(text, pos)
    return string.sub(text, pos, pos)
end

function Parser:new(code)
    local obj = {}
    obj.Code = code

    function obj:getCode()
        return obj.Code
    end

    function obj:Parse()
        TokenList = {}
        local curr_char
        local wasOp = true
        local tokenCount = 1

        for i = 0, string.len(obj.Code) do
            curr_char = getNextChar(obj.Code, i)

            if not (curr_char == '') and not (curr_char == ' ') and not (wasOp) then
                if type(curr_char) == "number" then
                    num = num + curr_char
                    wasOp = false
                elseif type(curr_char) == "string" and not wasOp then
                    TokenList[tokenCount] = Token.new(Token.NUMBER, num)
                    num = ''
                    tokenCount = tokenCount + 1
                    wasOp = true
                    if curr_char == "+" then
                        TokenList[tokenCount] = Token.new(Token.PLUS, curr_char)
                    elseif curr_char == "-" then
                        TokenList[tokenCount] = Token.new(Token.MINUS, curr_char)
                    elseif curr_char == "*" then
                        TokenList[tokenCount] = Token.new(Token.SUB, curr_char)
                    elseif curr_char == "/" then
                        TokenList[tokenCount] = Token.new(Token.DIV, curr_char)
                    end
                elseif type(curr_char) == "string" and wasOp then
                    error("Parse error: Ожидалось число, а обнаружен оператор: "..curr_char)
                end
            end
        end
        print(#TokenList)
        for i = 1, #TokenList do
            print("Token: "..TokenList[i].getType().." "..TokenList[i].getValue())
        end
        obj.TokenList = TokenList
    end

    function obj:Interp()
        local sum = 0
        local num = 0
        local op
        for i = 1, #obj.TokenList do
            if obj.TokenList[i].getType() == Token.NUMBER then
                num = num + tonumber(obj.TokenList.getValue())
            elseif obj.TokenList[i].getType() == Token.PLUS then

            end
        end
    end

    setmetatable(obj, self)
    self.__index = self; return obj
end

return Parser
