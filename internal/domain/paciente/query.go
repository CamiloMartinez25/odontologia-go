package paciente

var (
	QueryInsertPaciete = `INSERT INTO my_db.paciente(nombre, apellido, domicilio, dni, fecha_alta)
	VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta 
	FROM my_db.paciente`
	QueryDeletePaciente  = `DELETE FROM my_db.paciente WHERE id = ?`
	QueryGetPacienteById = `SELECT id, nombre, apellido, domicilio, dni, fecha_alta
	FROM my_db.paciente WHERE id = ?`
	QueryUpdatePaciente = `UPDATE my_db.paciente SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ?
	WHERE id = ?`
	QueryUpdateSubject = `UPDATE my_db.paciente SET `
)
