package main

type PtyRequestMsg struct {
  Term     string
  Columns  uint32
  Rows     uint32
  Width    uint32
  Height   uint32
  Modelist string
}
