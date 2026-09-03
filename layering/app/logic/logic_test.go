package logic

import (
	"styleguide/layering/httplayer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogic_RunA(t *testing.T) {
	// Тест бизнес-логики зависит от http слоя
	invalidRequest := httplayer.Request{
		Money: "A", // Не доменный тип данных
		Experiments: map[string]string{
			"A": "1",
		},
	}

	res, err := RunA(&invalidRequest)
	require.NoError(t, err)

	require.Equal(t, res, &httplayer.Response{
		Success: true,
		Status:  "A",
	})
}
