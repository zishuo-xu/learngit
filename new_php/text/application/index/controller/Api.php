<?php

namespace app\index\controller;

use think\Db;
use think\exception\HttpResponseException;
use think\Request;
use think\Response;
use think\response\Json;

class Api
{
	public function getUser() {
		if (Request::instance()->isGet()) {
			$user = Db::table('user')->field('uid,phone,email')->select();
			if (empty($user)) {
				$this->error('没有找到用户数据');
			} else {
				$this->success('获取成功', $user);
			}
		} else {
			$this->fail('请求失败');
		}
	}


	public function setInfo() {
		if (Request::instance()->isPut()) {
			$uid = Request::instance()->put('uid');
			$phone = Request::instance()->put('phone');
			$email = Request::instance()->put('email');
			if (empty($uid)) {
				$this->error('请输入uid');
			}
			if (empty($phone) && empty($email)) {
				$this->error('请至少修改一项');
			}
			$update = [];
			if ($phone) {
				$update['phone'] = $phone;
			}
			if ($email) {
				$update['email'] = $email;
			}
			$res = Db::table('user')->where('uid',$uid)->update($update);
			if ($res) {
				$this->success('修改成功');
			} else {
				$this->error('修改失败');
			}
		}
		$this->fail('请求失败');
	}






	public function success($msg, $data = []) {
		$status = 0;
		$result['status'] = $status;
		$result['msg'] = $msg;
		$result['data'] = $data;
		$response = Response::create($result, 'json', 0);
        throw new httpresponseexception($response);
	}

	public function error($msg, $data = []) {
		$status = -1;
		$result['status'] = $status;
		$result['msg'] = $msg;
		$result['data'] = $data;
		$response = Response::create($result, 'json', 0);
	}



	public function fail($msg = '请求失败') {
		$status = -1;
		$result['status'] = $status;
		$response = Response::create($result, 'json', -1);
	}


}
