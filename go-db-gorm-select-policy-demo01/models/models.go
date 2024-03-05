package models

import "time"

type Feedback struct {
  ID                         uint `gorm:"primarykey"`
  CreatedAt                  time.Time
  UserId                     string
  OverallExperience          int    // 1-5
  VoiceQuality               int    // 1-5
  ContextRelevancy           int    // 1-5
  CharacteristicsConsistency int    // 1-5
  TextFeedback               string // optional
}
