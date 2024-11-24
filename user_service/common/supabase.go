package common

import (
	"log"
	"user_service/utils"

	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

var AuthClient auth.Client

func InitSupabaseClient(){
    url := utils.AppConfig.Supabase.Url
	key := utils.AppConfig.Supabase.Key

	var err error
	SupabaseClient, err = supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("Failed to initialize Supabase client: %v", err)
	}

}

func GetSupabaseClient() *supabase.Client{
	// Initialize Supabase client if it's not already initialized
	if SupabaseClient == nil {
		utils.LoadConfig()
		InitSupabaseClient()
	}

	return SupabaseClient
}


func InitAuthClient(){
	project := utils.AppConfig.Supabase.Project
	key := utils.AppConfig.Supabase.Key

	AuthClient = auth.New(project, key)
}

func GetAuthClient() auth.Client{
	// Initialize Supabase client if it's not already initialized
	if AuthClient == nil {
		utils.LoadConfig()
		InitAuthClient()
	}

	return AuthClient
}