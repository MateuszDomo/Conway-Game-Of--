import sys

# i = Cur layer
# n = Number nodes to replicate
# ceil = Max layer
def dfs(i,n,ceil):

    if i == ceil:
        return 0

    sum = 0
    for j in range(n):
        sum += dfs(i+1,n,ceil) + 1

    return sum

def bfs(n,ceil):

    total = 0
    for i in range(1,ceil+1):
        total_cp = total
        total += n**i

        print(total-total_cp)


    return total

def main():
    if len(sys.argv) < 4:
            print("Not enough arguments")
            return

    n = int(sys.argv[1])
    ceil = int(sys.argv[2])


    if sys.argv[3] == "dfs":
        total = dfs(0,n,ceil)
    else:
        total = bfs(n,ceil)

    print()
    for i in range(1,10):
        print(dfs(0,n,i))
        print()

    print("total: ", str(total))

if __name__ == "__main__":
    main()



