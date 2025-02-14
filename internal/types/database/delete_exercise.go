package database

import "context"

// Интерфейс запроса на удаление упражнения
type DeleteExerciseReq interface {
	GetID() string
}

// Интерфейс ответа на удаление упражнения
type DeleteExerciseResp interface{}

// Структура, реализующая DeleteExerciseResp (пустая, так как ответ не требуется)
type deleteExerciseResp struct{}

// Метод удаления упражнения в `Database`
func (d *Database) DeleteExercise(ctx context.Context, req DeleteExerciseReq) (DeleteExerciseResp, error) {
	_, err := d.DB.ExecContext(ctx, `DELETE FROM exercises WHERE id = $1`, req.GetID())
	if err != nil {
		return nil, err
	}

	return deleteExerciseResp{}, nil
}
