import ast
data_file = "/home/meghshetty/Documents/Projects/Exercism/Python/two_fer.py"
data_terr = "/home/meghshetty/Documents/Projects/Exercism/Go_lang/astTester/terraform.tf"
f = open(data_file)
data =f.read()
code = "print('Hellowworld')"
tree = ast.parse(data)
print (ast.dump(tree, indent=4))