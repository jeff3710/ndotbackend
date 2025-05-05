package db

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
    // 测试数据1 - 用于创建新用户测试
    userParams1 := CreateUserParams{
        Username: "testuser_" + strconv.FormatInt(time.Now().UnixNano(), 10),
        Password: "hashed_password",
        Role:     "user",
        Active:   true,
    }

    // 测试数据2 - 用于更新用户测试
    userParams2 := CreateUserParams{
        Username: "testuser_" + strconv.FormatInt(time.Now().UnixNano(), 10),
        Password: "hashed_password",
        Role:     "user",
        Active:   true,
    }

    t.Run("创建新用户", func(t *testing.T) {
        err := testStore.CreateUser(context.Background(), userParams1)
        require.NoError(t, err)

        // 验证用户已创建
        user, err := testStore.GetUser(context.Background(), userParams1.Username)
        require.NoError(t, err)
        require.Equal(t, userParams1.Username, user.Username)
        
        // 清理测试数据
        defer func() {
            err := testStore.DeleteUser(context.Background(), userParams1.Username)
            require.NoError(t, err)
        }()
    })

    t.Run("更新已存在用户", func(t *testing.T) {
        // 先创建用户
        err := testStore.CreateUser(context.Background(), userParams2)
        require.NoError(t, err)

        // 更新用户信息
        updatedParams := userParams2
        updatedParams.Username = "updated_" + strconv.FormatInt(time.Now().UnixNano(), 10)

        err = testStore.CreateUser(context.Background(), updatedParams)
        require.NoError(t, err)

        // 验证用户已更新
        user, err := testStore.GetUser(context.Background(),updatedParams.Username)
        require.NoError(t, err)
        require.Equal(t, updatedParams.Username, user.Username)
        
        // 清理测试数据
        defer func() {
            err := testStore.DeleteUser(context.Background(), userParams2.Username)
            require.NoError(t, err)
        }()
    })
}