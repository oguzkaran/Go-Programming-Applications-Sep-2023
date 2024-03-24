package org.csystem.android.app.generator.random

import android.os.Bundle
import android.os.Handler
import android.os.Looper
import android.os.Message
import android.util.Log
import android.widget.ArrayAdapter
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.databinding.DataBindingUtil
import dagger.hilt.android.AndroidEntryPoint
import org.csystem.android.app.generator.random.databinding.ActivityMainBinding
import org.csystem.android.app.generator.random.global.what.WHAT_ANY_EXCEPTION
import org.csystem.android.app.generator.random.global.what.WHAT_GET_TEXT
import org.csystem.android.app.generator.random.global.what.WHAT_INVALID_VALUES
import org.csystem.android.app.generator.random.global.what.WHAT_IO_EXCEPTION
import org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo
import org.csystem.android.app.generator.random.viewmodel.data.ServerInfo
import org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel
import java.io.BufferedReader
import java.io.DataInputStream
import java.io.DataOutputStream
import java.io.IOException
import java.io.InputStreamReader
import java.lang.ref.WeakReference
import java.net.Socket
import java.nio.charset.StandardCharsets
import java.util.concurrent.ExecutorService
import javax.inject.Inject

@AndroidEntryPoint
class MainActivity : AppCompatActivity() {
    private lateinit var mBinding: ActivityMainBinding
    private lateinit var mHandler: Handler

    @Inject
    lateinit var threadPool: ExecutorService

    private class GetTextHandler(activity: MainActivity) : Handler(Looper.myLooper()!!) {
        private val mWeakReference = WeakReference(activity)

        override fun handleMessage(msg: Message)
        {
            val activity = mWeakReference.get()!!
            when (msg.what) {
                WHAT_GET_TEXT -> activity.handleGetText(msg.obj.toString())
                WHAT_INVALID_VALUES-> activity.handleInvalidValues(msg.obj.toString())
                WHAT_IO_EXCEPTION -> activity.handleIOException(msg.obj.toString())
                WHAT_ANY_EXCEPTION -> activity.handleAnyException(msg.obj.toString())
            }
        }
    }

    private fun handleGetText(text: String)
    {
        mBinding.adapter!!.add(text)
    }

    private fun handleInvalidValues(text: String)
    {
        Toast.makeText(this, "Invalid values:$text", Toast.LENGTH_LONG).show()
    }

    private fun handleIOException(message: String)
    {
        Log.d("con-io-ex", message)
        Toast.makeText(this, "IO problem occurred!...", Toast.LENGTH_LONG).show()
    }

    private fun handleAnyException(message: String)
    {
        Log.d("con-ex", message)
        Toast.makeText(this, "Problem occurred!...", Toast.LENGTH_LONG).show()
    }

    private fun generatorConnectionCallback(socket: Socket)
    {
        try {
            val dos = DataOutputStream(socket.getOutputStream())
            val dis = DataInputStream(socket.getInputStream())
            val br = BufferedReader(InputStreamReader(socket.getInputStream(), StandardCharsets.UTF_8))
            val count = mBinding.generateInfo!!.count

            dos.writeInt(count)
            dos.writeInt(mBinding.generateInfo!!.minimum)
            dos.writeInt(mBinding.generateInfo!!.maximum + 1)
            if (dis.readInt() == 0) {
                mHandler.sendMessage(mHandler.obtainMessage(WHAT_INVALID_VALUES, br.readLine()))
                return
            }

            (0..<count).forEach { _ -> mHandler.sendMessage(mHandler.obtainMessage(WHAT_GET_TEXT, br.readLine())) }
        }
        catch (ex: Throwable) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_ANY_EXCEPTION, ex.message))
        }
    }

    private fun configurationConnectionCallback(socket: Socket)
    {
        try {
            val br = BufferedReader(InputStreamReader(socket.getInputStream(), StandardCharsets.UTF_8))
            val configText = br.readLine().trim()

            runOnUiThread{Toast.makeText(this, configText, Toast.LENGTH_LONG).show()}
        }
        catch (ex: Throwable) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_ANY_EXCEPTION, ex.message))
        }
    }

    private fun textGeneratorCallback()
    {
        try {
            Socket(mBinding.serverInfo!!.host, mBinding.serverInfo!!.port).use{generatorConnectionCallback(it)}
        }
        catch (ex: IOException) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_IO_EXCEPTION, ex.message))
        }
        catch (ex: Throwable) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_GET_TEXT, ex.message))
        }
    }

    private fun configurationCallback()
    {
        try {
            Socket(mBinding.serverInfo!!.host, mBinding.serverInfo!!.port + 1).use{configurationConnectionCallback(it)}
        }
        catch (ex: IOException) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_IO_EXCEPTION, ex.message))
        }
        catch (ex: Throwable) {
            mHandler.sendMessage(mHandler.obtainMessage(WHAT_GET_TEXT, ex.message))
        }
    }

    private fun initViewModels()
    {
        mBinding.viewModel = MainActivityListenerViewModel(this)
        mBinding.generateInfo = GenerateInfo()
        mBinding.serverInfo = ServerInfo("192.168.1.107", 50527)
        mBinding.adapter = ArrayAdapter(this, android.R.layout.simple_list_item_1, ArrayList<String>())
        mBinding.enabled = true
    }

    private fun initBinding()
    {
        mBinding = DataBindingUtil.setContentView(this, R.layout.activity_main)
        initViewModels()
    }

    private fun initialize()
    {
        initBinding()
        mHandler = GetTextHandler(this)
    }

    override fun onCreate(savedInstanceState: Bundle?)
    {
        super.onCreate(savedInstanceState)
        initialize()
    }

    fun generateButtonClicked() = threadPool.execute{textGeneratorCallback()}

    fun configurationButtonClicked() = threadPool.execute{configurationCallback()}

    fun saveButtonClicked()
    {
        //TODO:
    }

    fun clearButtonClicked()
    {
        mBinding.adapter!!.clear()
    }
}