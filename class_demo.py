class Employee:
    def __init__(self, fname, lname, total_leaves, leaves_taken):
        self.first_name = fname
        self.last_name = lname
        self.total_leaves = total_leaves
        self.leaves_taken = leaves_taken

    def leaves_remaining(self):
        print(f"{self.first_name} {self.last_name} has {self.total_leaves - self.leaves_taken} leaves remaining")

e = Employee("Foo", "Bar", 30, 15)
e.leaves_remaining()
