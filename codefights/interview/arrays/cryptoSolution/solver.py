from re import sub

def solve(q):
  try:
    n = (i for i in q if i.isalpha()).next()
  except StopIteration:
    return q if eval(sub(r'(^|[^0-9])0+([1-9]+)', r'\1\2', q)) else False
  else:
    for i in (str(i) for i in range(10) if str(i) not in q):
      r = solve(q.replace(n,str(i)))
      if r:
        return r
    return False

if __name__ == "__main__":
  query = "SEND + 1ORE == 1ONEY"
  r = solve(query)
  print r if r else "No solution found."

# Other puzzles to try:
# query = "REASON == IT * IS + THERE"
# query = "MAD * MAN == ASYLUM"
# query = "THREE + THREE + ONE == SEVEN"
# query = "SEND + MORE == MONEY"
# query = "I + BB == ILL"
# query = "WHOSE + TEETH + ARE + AS == SWORDS"
# query = "BILL + WILLIAM + MONICA == CLINTON"
# query = "GREEN + ORANGE == COLORS"
# query = "PACIFIC + PACIFIC + PACIFIC == ATLANTIC"
# query = "CASSATT + RENOIR == PICASSO"
# query = "MANET + MATISSE + MIRO + MONET + RENOIR == ARTISTS"
# query = "COMPLEX + LAPLACE == CALCULUS"
