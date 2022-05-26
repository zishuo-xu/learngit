<?php
namespace app\index\controller;
use think\Controller;
use think\exception\HttpResponseException;
use think\Db;
use think\Request;
use think\response\Json;
use think\Session;

class Index extends Controller
{
    public function index()
    {
	    return view("index");
    }


    public function login(){
        //判断是否是否是post请求  如果不是就返回到登录界面 如果是就判断登录逻辑
	    if (Request::instance()->isPost()) {
            //接收用户名和密码
            $username = input('username');
            $password = input('password');
            //在数据库中查找有没有这个人
            $user = Db::table('user')->where('username', $username)->where('password', md5($password))->find();
            // 一般把用户信息存入session，记录登录状态  登录成功
            if ($user) {
                Session::set('name', $user['name']);
                $this->success('登录成功11', 'index');
            } else {
                //不存在 登录失败
                $this->error('登录失败');
            }
        }
        return view('login');
    }

    public function register(){
        //判断是否是否是post请求     如果不是就返回到注册界面 如果是就判断注册逻辑
	    if (Request::instance()->isPost()) {
		    $username = input('username');
		    $password = input('password');
		    $name = input('name');
		    $phone = input('phone');
		    $email = input('email');
		    if(empty($username))$this->error('请输入用户名');
		    if(empty($password))$this->error('请输入密码');
		    if(empty($name))$this->error('请输入姓名');
		    if(empty($phone))$this->error('请输入手机号');
		    if(empty($email))$this->error('请输入邮箱');
		    $user['username'] = $username;
		    $is_reg = Db::table('user')->where('username', $username)->find();
		    if ($is_reg) {
			    $this->error('用户已存在');
		    }
		    $user['password'] = md5($password);
		    $user['name'] = $name;
		    $user['phone'] = $phone;
		    $user['email'] = $email;
		    $result = Db::table('user')->insert($user);
		    if ($result) {
			    $this->success('注册成功','login');
		    } else {
			    $this->error('注册失败');
		    }
	    }
        return view('register');

    }

    public function logout(){
	    Session::set('name','');
	    return $this->success('退出成功','index');
    }


}
