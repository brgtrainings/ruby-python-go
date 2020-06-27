class Employee
  def initialize(fname, lname, total_leaves, leaves_taken)
    @first_name = fname
    @last_name = lname
    @total_leaves = total_leaves
    @leaves_taken = leaves_taken
  end

  def leaves_remaining()
    puts("#{@first_name} #{@last_name} has #{leaves_left} leaves remaining")
  end

  private

  def leaves_left
    @total_leaves - @leaves_taken
  end
end

e = Employee.new("Foo", "Bar", 30, 15)
e.leaves_remaining
