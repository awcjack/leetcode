/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportanceHelper(employeeHashmap map[int]*Employee, id int) int {
	importance := (*employeeHashmap[id]).Importance
	for _, employeeId := range (*employeeHashmap[id]).Subordinates {
		importance += getImportanceHelper(employeeHashmap, employeeId)
	}
	return importance
}

func getImportance(employees []*Employee, id int) int {
	employeeHashmap := make(map[int]*Employee)
	for _, employee := range employees {
		employeeHashmap[employee.Id] = employee
	}
	return getImportanceHelper(employeeHashmap, id)
}
