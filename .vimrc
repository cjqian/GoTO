set nocompatible              " be iMproved, required
filetype off                  " required

"runtime
execute pathogen#infect()

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')

" let Vundle manage Vundle, required
Plugin 'gmarik/Vundle.vim'
Plugin 'scrooloose/nerdtree'
Plugin 'fatih/vim-go'

nmap <silen><c-n> :NERDTreeToggle<CR> "map ctrl+n to nerdtree
call vundle#end()            " required
filetype plugin indent on    " required

"change colorscheme
syntax on
colorscheme molokai

"set highlight search
set hlsearch

"change font
set guifont=Menlo\ Regular:h12

"autoformat json
map ,j :%!python -m json.tool<CR>

  set tabstop=4                        " tabstops
  set shiftwidth=2                     " shift width
  "set expandtab                        " turn ^T to spaces
  set tabpagemax=15                    " only show 15 tabs

  if has('cmdline_info')
    set ruler                          " show the ruler
    set rulerformat=%30(%=\:b%n%y%m%r%w\ %l,%c%V\ %P%) " a ruler on steroids
    set showcmd                        " show partial commands in status line and
                                       " selected characters/lines in visual mode
  endif

  set nu

"swap window layouts
function! MarkWindowSwap()
    let g:markedWinNum = winnr()
endfunction

function! DoWindowSwap()
    "Mark destination
    let curNum = winnr()
    let curBuf = bufnr( "%" )
    exe g:markedWinNum . "wincmd w"
    "Switch to source and shuffle dest->source
    let markedBuf = bufnr( "%" )
    "Hide and open so that we aren't prompted and keep history
    exe 'hide buf' curBuf
    "Switch to dest and shuffle source->dest
    exe curNum . "wincmd w"
    "Hide and open so that we aren't prompted and keep history
    exe 'hide buf' markedBuf 
endfunction

nmap <silent> <leader>mw :call MarkWindowSwap()<CR>
nmap <silent> <leader>pw :call DoWindowSwap()<CR>

"set paste 
set paste
