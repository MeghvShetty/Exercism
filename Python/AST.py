import ast
data_file = "/home/meghshetty/Documents/Projects/Exercism/Python/two_fer.py"
data_terr = "/home/meghshetty/Documents/Projects/Exercism/Go_lang/astTester/terraform.tf"
f = open(data_terr)
data =f.read()
code = "print('Hellowworld')"
tree = ast.parse(code)
print (ast.dump(tree, indent=4))