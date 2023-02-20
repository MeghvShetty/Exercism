import ast
data_file = "/home/meghshetty/Documents/Projects/Exercism/Python/lasagna.py"
f = open(data_file)
data =f.read()
code = "print('Hellowworld')"
tree = ast.parse(code)
print (ast.dump(tree, indent=4))