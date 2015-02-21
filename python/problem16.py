result = 2**1000
total = 0
while result > 0:
  total += result % 10
  result = result / 10
print(total)
