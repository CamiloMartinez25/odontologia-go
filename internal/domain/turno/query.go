package turno

var (
	QueryInsertTurn = `INSERT INTO my_db.Turns(paciente,odontologo,fechaHora,descripcion)
	VALUES(?,?,?,?)`
	QueryGetAllTurns = `SELECT id, paciente,odontologo,fechaHora,descripcion 
	FROM my_db.Turns`
	QueryDeleteTurn  = `DELETE FROM my_db.Turns WHERE id = ?`
	QueryGetTurnById = `SELECT id, paciente,odontologo,fechaHora,descripcion
	FROM my_db.Turns WHERE id = ?`
	QueryGetTurnByPacienteId = `SELECT id, paciente,odontologo,fechaHora,descripcion
	FROM my_db.Turns WHERE paciente = ?`
	QueryUpdateTurn = `UPDATE my_db.Turns SET paciente = ?, odontologo = ?,  fechaHora = ?, descripcion = ?
	WHERE id = ?`
)
