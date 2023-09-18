package turno

var (
	QueryInsertTurn = `INSERT INTO my_db.turno(DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion)
	VALUES(?,?,?,?)`
	QueryInsertTurnByPaciente = `INSERT INTO my_db.turno(DNI_paciente,Matricula_odontologo)
	VALUES(?,?)`
	QueryGetAllTurns = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion 
	FROM my_db.turno`
	QueryDeleteTurn  = `DELETE FROM my_db.turno WHERE ID = ?`
	QueryGetTurnByID = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion
	FROM my_db.turno WHERE ID = ?`
	QueryGetTurnByPacienteID = `SELECT ID, DNI_paciente,Matricula_odontologo,Fecha_Hora,Descripcion
	FROM my_db.turno WHERE DNI_paciente = ?`
	QueryUpdateTurn = `UPDATE my_db.turno SET DNI_paciente = ?, Matricula_odontologo = ?,  Fecha_Hora = ?, Descripcion = ?
	WHERE ID = ?`
	QueryUpdateTurnoSubject = `UPDATE my_db.turno SET `

)
