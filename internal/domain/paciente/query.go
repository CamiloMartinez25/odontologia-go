package paciente

var (
	QueryInsertPaciete = `INSERT INTO my_db.pacientes(nombre, apellido, domicilio, dni, fecha_alta)
	VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta 
	FROM my_db.pacientes`
	QueryDeletePaciente  = `DELETE FROM my_db.pacientes WHERE id = ?`
	QueryGetPacienteById = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta
	FROM my_db.pacientes WHERE id = ?`
	QueryUpdatePaciente = `UPDATE my_db.pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ?
	WHERE id = ?`
)
