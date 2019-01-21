@inline function readparseinput(filename::String)
    arr = Vector{Vector{String}}()
    open(filename, "r") do fh
        for ln in eachline(fh)
            ln_parsed = split(ln, ",")
            push!(arr, ln_parsed)
        end
    end
    arr
end

mutable struct Townstats
    name::String
    transactions::Int
    maxpsm::Int
    avgpsm::Int
    minpsm::Int
end

function main()
    global arr = readparseinput("resale-flat-prices.csv")
    for l in arr
    end
end

function gettown(name::String, towns::Vector{Townstats})
    for t in towns
        if t.name == name
            return t
        end
    end
    return Townstats(name, 1, )
end
