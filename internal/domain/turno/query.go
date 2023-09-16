package turno

var (
	QueryInsertTurn = `INSERT INTO my_db.Turns(DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion)
	VALUES(?,?,?,?)`
	QueryInsertTurnByPaciente = `INSERT INTO my_db.Turns(DNI_paciente,Matricula_odontologo)
	VALUES(?,?)`
	QueryGetAllTurns = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion 
	FROM my_db.Turns`
	QueryDeleteTurn  = `DELETE FROM my_db.Turns WHERE ID = ?`
	QueryGetTurnByID = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion
	FROM my_db.Turns WHERE ID = ?`
	QueryGetTurnByPacienteID = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion
	FROM my_db.Turns WHERE DNI_paciente = ?`
	QueryUpdateTurn = `UPDATE my_db.Turns SET DNI_paciente = ?, Matricula_odontologo = ?,  Fecha_Hora = ?, Descripcion = ?
	WHERE ID = ?`
)
