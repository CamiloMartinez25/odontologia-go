package paciente

var (
	QueryInsertPaciente = `INSERT INTO my_db.paciente(Nombre, Apellido, Domicilio, DNI, Fecha_Alta)
	VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT ID, Nombre, Apellido, Domicilio, DNI, Fecha_Alta 
	FROM my_db.paciente`
	QueryDeletePaciente  = `DELETE FROM my_db.paciente WHERE ID = ?`
	QueryGetPacienteById = `SELECT ID, Nombre, Apellido, Domicilio, DNI, Fecha_Alta
	FROM my_db.paciente WHERE ID = ?`
	QueryUpdatePaciente = `UPDATE my_db.paciente SET Nombre = ?, Apellido = ?, Domicilio = ?, DNI = ?, Fecha_Alta = ?
	WHERE ID = ?`
	QueryUpdateSubject = `UPDATE my_db.paciente SET `
)
