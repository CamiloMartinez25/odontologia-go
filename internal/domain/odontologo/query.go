package odontologo

var (
	QueryInsertOdontologo       = `INSERT INTO my_db.odontologos(Nombre, Apellido, Matricula) VALUES(?,?,?)`
	QueryGetAllOdontologos      = `SELECT ID, Nombre, Apellido, Matricula FROM my_db.odontologos`
	QueryDeleteOdontologo       = `DELETE FROM my_db.odontologos WHERE ID = ?`
	QueryGetOdontologoById      = `SELECT ID, Nombre, Apellido, Matricula FROM my_db.odontologos WHERE ID = ?`
	QueryUpdateOdontologo       = `UPDATE my_db.odontologos SET Nombre = ?, Apellido = ?, Matricula = ? WHERE ID = ?`
	QueryUpdateOdontologoSubject = `UPDATE my_db.odontologos SET `
)
